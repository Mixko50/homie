//
//  GroupList.swift
//  Homie
//
//  Created by Mixko on 10/6/22.
//

import SwiftUI

struct GroupList: View {
    var body: some View {
        let sampleTime = "11-10-2022 20:39"
        VStack(alignment: .trailing) {
            HStack {
                HStack{
                    Image(systemName: "person.3.fill").foregroundColor(Color.white).frame(width: 40, height: 40, alignment: .center).background(Color.pink).cornerRadius(10)
                    VStack (alignment: .leading) {
                        Text("Mixko's home").font(.system(size: 16, weight: .bold, design: .rounded))
                        Text("Created at : ").font(.system(size: 14, weight: .medium)) +  Text(sampleTime).font(.system(size: 14, weight: .light))
                    }
                }
            }
        }
    }
}

struct GroupList_Previews: PreviewProvider {
    static var previews: some View {
        GroupList()
    }
}
