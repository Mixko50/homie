//
//  StateBox.swift
//  Homie
//
//  Created by Mixko on 10/5/22.
//

import SwiftUI

struct StateBox: View {
    var body: some View {
        let sampleText = "on"
        let sampleTime = "11-10-2022 20:39"
        VStack(alignment: .trailing) {
            HStack {
                HStack{
                    Image(systemName: "externaldrive.connected.to.line.below.fill").foregroundColor(Color.white).frame(width: 40, height: 40, alignment: .center).background(Color.pink).cornerRadius(10)
                    VStack (alignment: .leading) {
                        Text("Air Conditioner").font(.system(size: 16, weight: .medium, design: .rounded))
                        Text(sampleTime).font(.system(size: 14, weight: .light))
                    }
                }
                Spacer()
                Text(sampleText).font(.system(size: 18, weight: .bold, design: .rounded)).foregroundColor(sampleText == "on" ? Color.green : Color.red)
            }
        }
    }
}

struct StateBox_Previews: PreviewProvider {
    static var previews: some View {
        StateBox()
    }
}
